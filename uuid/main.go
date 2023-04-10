package main

import (
    "log"
    "os"
    "encoding/json"

    uuid "github.com/google/uuid"
    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func generate_uuid() (string, error) {
    id, err := uuid.NewRandom()
    if err != nil {
        return "", err
    }
    return id.String(), nil
}

func main() {
    n := maelstrom.NewNode()
    
    n.Handle("generate", func(msg maelstrom.Message) error {
        var body map[string]any
        if err := json.Unmarshal(msg.Body, &body); err != nil {
            return err
        }

        body["type"] = "generate_ok"
        body["id"], _ = generate_uuid()

        return n.Reply(msg, body)
    })

    if err := n.Run(); err != nil {
        log.Printf("ERROR: %s", err)
        os.Exit(1)
    }
}
