# daybreakWidescreen

A Higurashi Daybreak Kai patcher/mod to get true widescreen, supporting various resolutions.

## How to Use

## Prerequisites

1. Ensure you have the game's folder ready.
2. Verify that the required files (`DaybreakDX.exe` and `config.dat`) are present in the game's directory.

## Using the GUI Mode

1. Extract the zip contents into the game's folder.
2. Run `wideMod.exe` found in the [releases](https://github.com/Vmarcelo49/daybreakWidescreen/releases).
3. Follow the on-screen instructions to set your desired resolution and fullscreen mode.

## Using the CLI Mode

These are the avaliable commands:

   ```bash
   wideMod.exe --resolution=<width>x<height>
   wideMod.exe --revert
   wideMod.exe --name <yourName>
   wideMod.exe --shadow
   wideMod.exe --outline
   wideMod.exe --htexture
   
   ```

   Replace `<width>x<height>` with your desired resolution (e.g., `1920x1080`) and `<true|false>` with your fullscreen preference.

![image](https://github.com/user-attachments/assets/486da64f-ccfd-443d-9ca3-f2e666ecfeb8)

![image](https://github.com/user-attachments/assets/da1af8fb-5cc9-4061-9467-dcc02a280de1)

![image](https://github.com/user-attachments/assets/7b33b3fa-2baf-4e21-a6a0-11453f68f40e)

## Thanks to

- Marchizu, for helping to find the offsets in the .exe
- Olof-AV, for creating the fix
- NigelTicknor, for creating his launcher, which inspired some of the features used here.
