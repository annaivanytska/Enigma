# Enigma

### Overview
This Go program simulates the Enigma encryption machine used during World War II. It includes functions for setting up rotors, handling text input, performing encryption, and managing rotor rotations.

### Key Components

1. **Struct Definition**:
   ```go
   type Rotors struct {
       Table  []int // Mapping for letters
       Flaga  int   // Rotor turnover position
       CurPos int   // Current position
       Turn   bool  // Whether the rotor should turn
   }
   ```

2. **Input Functions**:
   - `InputRotors()`: Prompts the user to select three rotors and their initial positions.
   - `InputText()`: Reads input from the console.

3. **Rotor Setup**:
   - `MakeRotors()`: Configures the selected rotors based on user input and sets their initial positions.

4. **Encryption Path**:
   - `ForwardPath()`: Processes the letter through the rotors in the forward direction.
   - `Miror()`: Simulates the reflector by mapping letters to their mirrored counterparts.
   - `ReturnPath()`: Processes the letter through the rotors in the reverse direction.

5. **Rotor Rotation**:
   - `Rotation()`: Manages the rotation of the rotors after each letter is processed.

6. **Character Mapping**:
   - `FindByte()`: Converts a character to its corresponding rotor position.
   - `FindLetter()`: Converts a rotor position back to a character.

7. **Encryption Process**:
   - `Enigma()`: Main encryption function that processes the input text, applies the rotor transformations, and outputs the encrypted text.

8. **File Handling**:
   - `ReadText()`: Reads text from a file if the user chooses to encrypt text from a file.

9. **User Interface**:
   - `Menu()`: Displays a menu for the user to choose between inputting text directly or reading from a file.
   - `main()`: Entry point of the program that calls the menu, reads the input, and starts the encryption process.

### Main Function Flow
1. The user is prompted to choose input type (text or file).
2. Depending on the input type, the text is either read from the console or a file.
3. The user selects the rotors and their initial positions.
4. The `Enigma` function processes the text, applying the encryption mechanism.
5. The encrypted text is printed as output.

This program mimics the workings of the historical Enigma machine, allowing users to encrypt messages using a configurable rotor-based system.
