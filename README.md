# goroutines

Projeto de estudos para entender como funcionam as goroutines em Go.

## Descrição

Este projeto foi criado com o objetivo de explorar e aprender sobre as goroutines, mecanismos concorrentes do Go. Aqui, experimentamos a criação de consumidores e produtores que se comunicam através de canais para executar tarefas de forma concorrente.

A task é genérica, executando apenas um `time.Sleep(<duration>)`, no qual a duração é dada por um número randômico entre 0 e 5 segundo calculado em tempo de execução da task, apenas para simular um processamento demorado.

## Como executar

1. Certifique-se de ter o Go instalado.
2. Navegue até a pasta raiz do projeto:
   ```bash
   cd .../goroutines
   ```
3. Execute o comando abaixo, substituindo `<n_workers>` pelo número desejado de consumidores:
    ```bash
    go run cmd/main.go <n_workers>
    ```
4. O programa irá iniciar o produtor e os consumidores, processando tarefas conforme definido.

## Observações

- Este projeto é destinado a experimentações e aprendizado.
- Explore o código fonte em `internal/entities` para entender a implementação dos consumidores e do produtor.