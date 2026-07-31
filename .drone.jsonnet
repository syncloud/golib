[{
    kind: "pipeline",
    type: "docker",
    trigger: {
        event: [ "push" ]
    },
    steps: [
        {
            name: "test",
            image: "golang:1.20",
            commands: [
                "go test ./...",
                "go build ./...",
            ]
        },
    ],
}]