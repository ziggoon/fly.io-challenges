package main

import (
    "encoding/json"
    "container/list"
    "log"
    "os"

    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
    n := maelstrom.NewNode()
    values_received := list.New()
    
    n.Handle("broadcast", func(msg maelstrom.Message) error {
        body := map[string]any{
            "type": "broadcast_ok",
        }
        values_received.PushBack(body["message"])
        return n.Reply(msg, body)
    })

    n.Handle("read", func(msg maelstrom.Message) error {
        body := map[string]any{
            "messages": values_received,
            "type": "read_ok",
        }
        return n.Reply(msg, body)
    })

    n.Handle("topology", func(msg maelstrom.Message) error {
        var body map[string] any
        if err := json.Unmarshal(msg.Body, &body); err !=
        nil {
            return err
        }
        
        body["type"] = "topology_ok"
        delete(body, "topology")
        return n.Reply(msg, body)
    })

    if err := n.Run(); err != nil {
        log.Printf("ERROR: %s", err)
        os.Exit(1)
    }
}
