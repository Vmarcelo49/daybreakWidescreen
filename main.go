package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gonutz/wui/v2"
)

var resolutions = []string{"1152x648", "1280x720", "1366x768", "1600x900", "1920x1080", "2560x1440", "3840x2160"}

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

	// valores padrão de escalamento
	for _, item := range hudItems {
		item.setDefaultScaledValues(float32(width), float32(height))
	}

	for _, item := range hudItems {
		item.apply(file)
	}

	return nil
}

func revertToOriginalEXE() {
	// verifica se existe um backup
	_, err := os.Stat("./backup/DaybreakDX.exe")
	if os.IsNotExist(err) {
		throwErrorMessageWindow("Backup not found")
	}
	// apaga o arquivo atual
	err = os.Remove("./DaybreakDX.exe")
	if err != nil {
		throwErrorMessageWindow("Error while deleting current .exe file" + err.Error())
	}
	// copia o backup para o local original
	err = copyFile("./backup/DaybreakDX.exe", "./DaybreakDX.exe")
	if err != nil {
		throwErrorMessageWindow("Error while copying backup file" + err.Error())
	}
}

func copyFile(src, dst string) error {
	// Abre o arquivo de origem
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Cria o arquivo de destino
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copia o conteúdo
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Força o buffer a gravar no disco
	return destFile.Sync()
}

func createBackup() error {
	backupPath := "./backup/DaybreakDX.exe"
	err := os.MkdirAll("backup", os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		err := copyFile("./DaybreakDX.exe", backupPath)
		if err != nil {
			return err
		}
		fmt.Println("Arquivo copiado com sucesso!")
	} else {
		fmt.Println("Arquivo já existe, sem necessidade de cópia.")
	}
	return nil

}

func parseResolution(resolution string) (uint16, uint16) {
	strValues := strings.Split(resolution, "x")
	intWidth, _ := strconv.Atoi(strValues[0])
	intHeight, _ := strconv.Atoi(strValues[1])
	return uint16(intWidth), uint16(intHeight)
}

func patchAndSave(resolutionDropdown string, fullscreen bool) {
	// Primeiro fazendo um backup do arquivo
	err := createBackup()
	if err != nil {
		fmt.Println("Erro ao fazer backup do arquivo:", err)
	}

	// Convertendo a string para uint16
	width, height := parseResolution(resolutionDropdown)

	// Modificar o arquivo executável (especificar o caminho)
	err = modifyExecutable("DaybreakDX.exe", width, height)
	if err != nil {
		throwErrorMessageWindow("Error while modifying .exe file" + err.Error())
	} else {
		fmt.Println("Modificações concluídas com sucesso!")
	}
}

func verifyRequiredFiles() {
	requiredFiles := []string{"./DaybreakDX.exe", "./config.dat"}
	for _, file := range requiredFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			showError(file + " not found in current directory")
			if file == "./config.dat"{
				showError("open config in the game to generate the config.dat")
			}
		}
	}
}

func main() {
	verifyRequiredFiles()
	window := wui.NewWindow()
	configWindow(window)

	window.Show()
}
