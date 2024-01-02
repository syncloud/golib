[{
    kind: "pipeline",
    type: "docker",
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