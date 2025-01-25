package save

import "os"

func SaveToFile(filename, value string) error {
    err := os.WriteFile(filename, []byte(value), os.FileMode(os.O_WRONLY))

    if err != nil {
        return err
    }

    return nil
}
