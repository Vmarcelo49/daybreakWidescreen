package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	Width      uint16
	Height     uint16
	Fullscreen bool
}

func main() {
	// Cria a aplicação Fyne
	myApp := app.New()
	w := myApp.NewWindow("Configuration Tool")
	w.Resize(fyne.NewSize(350, 250))

	// Opções de resolução
	resolutions := []string{"1152x648", "1280x720", "1366x768", "1600x900", "1920x1080", "2560x1440", "3840x2160"}

	config, err := readConfigFile("windowConfig.ini")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		config = Config{ // Valores padrão se não houver arquivo de configuração
			Width:      1280,
			Height:     720,
			Fullscreen: false,
		}
	}

	// Criando os dropdowns
	resolutionDropdown := widget.NewSelect(resolutions, func(selected string) {})
	resolutionDropdown.Selected = fmt.Sprintf("%dx%d", config.Width, config.Height)

	fullscreenCheckbox := widget.NewCheck("Fullscreen", func(checked bool) {
		config.Fullscreen = checked
	})
	fullscreenCheckbox.SetChecked(config.Fullscreen)

	// Botão de confirmação
	confirmButton := widget.NewButton("Patch and Save", func() {
		patchAndSave(resolutionDropdown, fullscreenCheckbox.Checked)
		w.Close()

	})
	confirmAndLaunchButton := widget.NewButton("Patch, Save and Launch", func() {
		patchAndSave(resolutionDropdown, fullscreenCheckbox.Checked)
		w.Close()
		// Executar o jogo
		cmd := exec.Command("./DaybreakFixLoader.exe")
		cmd.Start()
	})

	// Layout da janela
	content := container.NewVBox(
		widget.NewLabel("Resolution:"),
		resolutionDropdown,
		fullscreenCheckbox,
		confirmButton,
		confirmAndLaunchButton,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

func patchAndSave(resolutionDropdown *widget.Select, fullscreen bool) {
	selectedResolution := resolutionDropdown.Selected

	// Converter a resolução selecionada
	var width, height uint16
	switch selectedResolution {
	case "1152x648":
		width, height = 1152, 648
	case "1280x720":
		width, height = 1280, 720
	case "1366x768":
		width, height = 1366, 768
	case "1600x900":
		width, height = 1600, 900
	case "1920x1080":
		width, height = 1920, 1080
	case "2560x1440":
		width, height = 2560, 1440
	case "3840x2160":
		width, height = 3840, 2160
	}

	// Modificar o arquivo executável (especificar o caminho)
	err := modifyExecutable("DaybreakDX.exe", width, height)
	if err != nil {
		fmt.Println("Erro ao modificar o arquivo:", err)
	} else {
		fmt.Println("Modificações concluídas com sucesso!")
	}
	err = writeConfigFile(width, height, fullscreen)
	if err != nil {
		fmt.Println("Erro ao escrever o arquivo de configuração:", err)
	} else {
		fmt.Println("Arquivo de configuração escrito com sucesso!")
	}
}

func readConfigFile(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "width=") {
			widthStr := strings.TrimPrefix(line, "width=")
			width, err := strconv.Atoi(widthStr)
			if err != nil {
				return config, fmt.Errorf("erro ao converter largura: %w", err)
			}
			config.Width = uint16(width)
		} else if strings.HasPrefix(line, "height=") {
			heightStr := strings.TrimPrefix(line, "height=")
			height, err := strconv.Atoi(heightStr)
			if err != nil {
				return config, fmt.Errorf("erro ao converter altura: %w", err)
			}
			config.Height = uint16(height)
		} else if strings.HasPrefix(line, "fullscreen=") {
			fullscreenStr := strings.TrimPrefix(line, "fullscreen=")
			fullscreen, err := strconv.ParseBool(fullscreenStr)
			if err != nil {
				return config, fmt.Errorf("erro ao converter fullscreen: %w", err)
			}
			config.Fullscreen = fullscreen
		}
	}

	if err := scanner.Err(); err != nil {
		return config, err
	}

	return config, nil
}

func writeConfigFile(width, height uint16, fullscreen bool) error {
	// verifica se windowConfig.ini existe
	if _, err := os.Stat("windowConfig.ini"); os.IsNotExist(err) {
		// cria o arquivo
		file, err := os.Create("windowConfig.ini")
		if err != nil {
			return fmt.Errorf("erro ao criar o arquivo: %w", err)
		}
		defer file.Close()
	}
	// sobre-escreve o arquivo existente
	file, err := os.OpenFile("windowConfig.ini", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	// escreve as configurações
	file.WriteString("[Window]\n")
	file.WriteString(fmt.Sprintf("width=%d\n", width))
	file.WriteString(fmt.Sprintf("height=%d\n", height))
	file.WriteString(fmt.Sprintf("fullscreen=%t\n", fullscreen))

	return nil
}

func writeFloat32BytesToFile(file *os.File, offset int64, newValue float32) error {
	file.Seek(offset, 0)
	err := binary.Write(file, binary.LittleEndian, newValue)
	return err
}

// Função para modificar o arquivo executável
func modifyExecutable(fileName string, width, height uint16) error {
	// Abre o arquivo para leitura e escrita
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	file.Seek(0x2338A0, 0)                     // Modifica a aspect ratio
	file.Write([]byte{0x39, 0x8E, 0xE3, 0x3F}) // 16:9
	// uint16 SECTION Rendering stuff

	file.Seek(0x2397FF, 0) // offset para a largura de renderização
	binary.Write(file, binary.LittleEndian, width)

	// Modifica a altura de renderização
	file.Seek(0x2397FA, 0) // offset para a altura de renderização
	binary.Write(file, binary.LittleEndian, height)

	// Modifica a largura da janela
	file.Seek(0x24D2B3, 0) // offset para a largura da janela
	binary.Write(file, binary.LittleEndian, width)

	// Modifica a altura da janela
	file.Seek(0x24D2D9, 0) // offset para a altura da janela
	binary.Write(file, binary.LittleEndian, height)

	file.Seek(0x239793, 0)
	binary.Write(file, binary.LittleEndian, width)

	file.Seek(0x23978E, 0)
	binary.Write(file, binary.LittleEndian, height)

	// float32 SECTION

	halfWidth := float32(width / 2)
	halfHeight := float32(height / 2)

	writeFloat32BytesToFile(file, 0x2A16D4, halfWidth)  //alvo ingame x e parte do char select
	writeFloat32BytesToFile(file, 0x2A1720, halfHeight) //alvo ingame y

	writeFloat32BytesToFile(file, 0xCC60, float32(width))  // X screen flash
	writeFloat32BytesToFile(file, 0xCC5B, float32(height)) // Y screen flash

	// Char select
	writeFloat32BytesToFile(file, 0xCCC1, float32(width))  // X screen effect?
	writeFloat32BytesToFile(file, 0xCCBC, float32(height)) // Y screen effect?

	writeFloat32BytesToFile(file, 0x2A16D8, float32(halfWidth)) // xPortrait
	writeFloat32BytesToFile(file, 0x8B4A, float32(halfHeight))  // yPortrait

	// HUD elements
	writeFloat32BytesToFile(file, 0x2A4928, float32(width))  // burst x
	writeFloat32BytesToFile(file, 0x2A3B38, float32(height)) // burst y

	writeFloat32BytesToFile(file, 0x23A03C, float32(100)) // fps counterx

	writeFloat32BytesToFile(file, 0x11599, float32(halfWidth)) // menu items x
	writeFloat32BytesToFile(file, 0x11545, float32(halfWidth)) // menu logo x

	hudItems := getHudValues(float32(width))

	//setting proportional values
	for _, item := range hudItems {
		item.setDefaultScaledValues(float32(width), float32(height))
	}

	//setting custom values
	for _, item := range hudItems {
		item.apply(file)
	}

	return nil
}
