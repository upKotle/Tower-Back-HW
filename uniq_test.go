package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestUniqWithoutParams(t *testing.T) {

	input := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	expectedOutput := `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
`

	reader, writer, _ := os.Pipe()
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	io.WriteString(writer, input)
	writer.Close()

	os.Stdin = reader

	outputReader, outputWriter, _ := os.Pipe()

	os.Stdout = outputWriter

	main()

	outputWriter.Close()

	var outputBuffer bytes.Buffer
	io.Copy(&outputBuffer, outputReader)

	actualOutput := outputBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод:\n%s\nФактический вывод:\n%s", expectedOutput, actualOutput)
	}
}
