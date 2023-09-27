# Concorrência

## **Concorrência x Paralelismo**

Paralelismo é quando você está fazendo duas coisas, exatamente ao mesmo tempo, e isso só é possível, se você tiver dois processadores. Fora isso, considere como sendo concorrência o ato de intercalar e administrar vários processos ao mesmo tempo e em um único processador físico.

## **Goroutine**

Thread -> O conceito é abrir uma trilha de execução independente do fluxo atual. No caso do Go, não criamos threads, mas sim Goroutines.

A Go routine é uma função que será executada de forma assíncrona.

```go
go anyFunc()
```

### **Channel**

O channel é um tipo de dados, assim como Boolean, Int... ele é utilizado como uma forma de "sincronizar" e comunicar as Goroutines, fazendo com que você consiga aguardar e unir informações/dados.

O canal sem buffer gera uma operação bloqueante que precisa ser aguardada (Ex: await);

```go
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)
```

## **Multiplexar**

...

## **Estrutura de Controle**

### Select

Estrutura de controle utilizada para trabalhar com concorrência.


## **Perguntas**

1 - O que é concorrência em Go e por que é importante?

Qual é a diferença entre concorrência e paralelismo em Go?

Como criar uma goroutine em Go?

O que é um canal (channel) em Go e como ele é usado para facilitar a comunicação entre goroutines?

Qual é a diferença entre um canal não bufferizado e um canal bufferizado em Go?

O que é um "mutex" em Go e qual é o seu papel na sincronização de acesso a recursos compartilhados?

Explique o conceito de "goroutine leakage" e como evitá-lo.

O que é "race condition" e como Go ajuda a prevenir esses problemas?

Como você pode usar a primitiva sync.WaitGroup em Go para esperar que um grupo de goroutines seja concluído?

Quais são as principais ferramentas e pacotes em Go para auxiliar na detecção e resolução de problemas de concorrência?