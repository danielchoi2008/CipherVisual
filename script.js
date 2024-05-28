function visualizeCipher() {
    const plaintext = document.getElementById('plaintext').value;
    const shift = parseInt(document.getElementById('shift').value);
    const visualization = document.getElementById('visualization');
    visualization.innerHTML = '';

    if (!plaintext) {
        alert('Please enter some plaintext');
        return;
    }

    const steps = [];

    // Step 1: Plaintext input
    steps.push(`Plaintext: ${plaintext}`);

    // Step 2: Cipher method and shift value
    steps.push(`Cipher Method: Caesar Cipher, Shift Value: ${shift}`);

    // Step 3: Convert plaintext to individual characters
    const characters = plaintext.split('');
    steps.push(`Characters: ${JSON.stringify(characters)}`);

    // Step 4: Shift each character
    const shiftedChars = [];
    for (const char of characters) {
        // Character must match regex
        if (char.match(/[a-z]/i)) {
            const charCode = char.charCodeAt(0);
            const base = charCode >= 65 && charCode <= 90 ? 65 : 97;
            const originalIndex = charCode - base;
            const shiftedIndex = (originalIndex + shift) % 26;
            const shiftedChar = String.fromCharCode(base + shiftedIndex);
            shiftedChars.push(shiftedChar);

            steps.push(`Search: Match Found! Index of character '${char}' was found on the alphabet at position ${originalIndex}!`);
            steps.push(`Shift: Character '${char}' shifted by ${shift} becomes '${shiftedChar}' at position ${shiftedIndex}!`);
        } else {
            shiftedChars.push(char);
            steps.push(`Non-alphabetic character: '${char}' remains unchanged`);
        }
    }

    // Step 5: Combine shifted characters to form ciphertext
    const ciphertext = shiftedChars.join('');
    steps.push(`Ciphertext: ${ciphertext}`);

    // Display steps
    for (const [index, step] of steps.entries()) {
        const stepDiv = document.createElement('div');
        stepDiv.className = 'step';
        stepDiv.innerText = `Step ${index + 1}: ${step}`;
        visualization.appendChild(stepDiv);
    }
}
