# Day 04 - Go Boot camp

## Сладость!

## Contents

1. [Глава I](#chapter-i) \
    1.1. [Основные правила](#general-rules)
2. [Глава II](#chapter-ii) \
    2.1. [Правила дня](#rules-of-the-day)
3. [Глава III](#chapter-iii) \
    3.1. [Введение](#intro)
4. [Глава IV](#chapter-iv) \
    4.1. [Упражнение 00: Поимка удачи](#exercise-00-catching-the-fortune)
5. [Глава V](#chapter-v) \
    5.1. [Упражнение 01: Закон и порядок](#exercise-01-law-and-order)
6. [Глава VI](#chapter-vi) \
    6.1. [Упражнение 02: Старая корова](#exercise-02-old-cow)
7. [Глава VII](#chapter-vii) \
    7.1. [Reading](#reading)


<h2 id="chapter-i" >Глава I</h2>
<h2 id="general-rules" >Основные правила</h2>

* Твоя программа не должна закрываться неожиданно (выдавая ошибку при корректном вводе). Если это произойдет, твой проект будет считаться неработаспособным и получит 0 во время оценки.
* Мы рекомендуем тебе писать тесты для твоего проекта, даже если если они и не оцениваются. Это даст тебе возможность легко тестировать твою работу и работу твоих пиров. Ты убедишься что тесты очень полезны, во время защиты. Во время защиты ты свободен использовать свои тесты и/или тесты пира которого ты проверяешь.
* Отправляй свою работу в нужный git репозиторий. Работа будет оцениваться только из git репозитория.
* Если твой код использует сторонние зависимости, следует использовать [Go Modules](https://go.dev/blog/using-go-modules) для управления ими.

<h2 id="chapter-ii" >Глава II</h2>
<h2 id="rules-of-the-day" >Правила дня</h2>


* Пиши код только в `*.go` файлах и (в случае стронних зависимостей) `go.mod` + `go.sum`
* Твой код для этого задания должен собираться с использовния простого `go build`
* Даже если не потребуется изменять С код, вам все равно придется закомментировать `main()` функцию, иначе твоя программа не скомпилируется (по причине двух точек входа в программу)


<h2 id="chapter-iii" >Глава III</h2>
<h2 id="intro" >Введение</h2>

Мистер роджерс очень грустный. Он сидит в своем офисе и болтает "Весь мой бизнес..Как я могу теперь сделать людей счастливыми?".

Эта история стара как мир. Этот ваш новый клиент открыл новый бизнес по продаже конфет по всему этому грязному городку. По началу все было прекрасно - несколько вендинговых автоматов, 5 вкуснейший типов конфет и очереди из детей упрашивающий своих родителей купить немного для них. Затем как гром среди ясного неба, когда кто-то вломился в дата-центр и украл сервер отвечающий за обработку заказов конфет. Не только это пошло не так, старый разработчик пропал, тоже! Совпадение? Вы так не думаете. 

Вы налили мистеру Роджерсу стакан хорошего старого бурбона и начали задавать вопросы пытаясь получить побольше деталей.

"Well, I don't know much. All our vending machines were selling the same set of candy, you know, here they are on the brochure" - he gives you the piece of colorful paper advertising five new amazing tastes:

"Ну, я не знаю многого. Все наши вендинговые автоматы продавали один и тот же набор конфет, ты знаешь, вот они все на брошюре" - он дает вам кусочек цветастой бумаги рекламирующей пять новых невероятных вкусов:
__EN__
<table>
  <tr>
    <th>Name</th>
    <th>Price</th>
  </tr>
  <tr>
    <td>Cool Eskimo</td>
    <td>10 cents</td>
  </tr>
  <tr>
    <td>Apricot Aardvark</td>
    <td>15 cents</td>
  </tr>
  <tr>
    <td>Natural Tiger</td>
    <td>17 cents</td>
  </tr>
  <tr>
    <td>Dazzling 	Elderberry</td>
    <td>21 cents</td>
  </tr>
  <tr>
    <td>Yellow Rambutan</td>
    <td>23 cents</td>
  </tr>
</table>

__RU__

<table>
  <tr>
    <th>Название</th>
    <th>Цена</th>
  </tr>
  <tr>
    <td>Крутое Эскимо</td>
    <td>10 центов</td>
  </tr>
  <tr>
    <td>Абрикосовый трубкозуб</td>
    <td>15 центов</td>
  </tr>
  <tr>
    <td>Натуральный тигр</td>
    <td>17 центов</td>
  </tr>
  <tr>
    <td>Ослепительная бузина</td>
    <td>21 цент</td>
  </tr>
  <tr>
    <td>Желтый рамбутан</td>
    <td>23 цента</td>
  </tr>
</table>

"Какие у них странно звучащие названия" - ты говоришь - "Как люди должны запоминать эти вещи?"
"Он ну это просто" - ответил Роджерс - "Мы используем аббревиатуры везде, включая наш исходный код, так оно и получается CE, AA, NT и так далее.."

Он всхлипнул.

"Но имеет ли это значение сейчас? Мой бизнес разваливается в любом случае, все это просто какой-то нонсенс!" 

"Пожалуйста, сфокусируйтесь, мистер Роджерс" - вы должно быть видели ребят которые вели себя так много раз, это место не просто так названо "Гофери ПиАй" - "Есть какие-нибудь детали которые вы еще не отметили?" 

"Вы правы! Я совсем забыл!" - он вытащил кусок бумаги из своего кармана и протянул его. - "Вор оставил записку!"

Вы посмотрели на текст написанный маркером на одной стороне: "Я не могу больше есть конфеты!". Это не дало много информации. Затем вы перевернули листок и..

"Окэй мистер Роджерс. Хорошая новость, я теперь знаю наверняка это был ваш бывший работник кто украл сервер. Но не только это! Что-то мне говорить что я могу помочь восстановить ваш бизнес тоже." 

<h2 id="chapter-iv" >Глава IV</h2>
<h3 id="ex00">Упражнение 00: Поимка удачи</h3>

Перевернув, вор использовал первый кусок бумаг на своем рабочем месте и по счастливому случаю это была спецификация для протокола между вендинговым автоматом и сервером. Он выглядии следующим образом:

```yaml
---
swagger: '2.0'
info:
  version: 1.0.0
  title: Candy Server
paths:
  /buy_candy:
    post:
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - in: body
          name: order
          description: summary of the candy order
          schema:
            type: object
            required:
              - money
              - candyType
              - candyCount
            properties:
              money:
                description: amount of money put into vending machine
                type: integer
              candyType:
                description: kind of candy
                type: string
              candyCount:
                description: number of candy
                type: integer
      operationId: buyCandy
      responses:
        201:
          description: purchase succesful
          schema:
              type: object
              properties:
                thanks:
                  type: string
                change:
                  type: integer
        400:
          description: some error in input data
          schema:
              type: object
              properties:
                error:
                  type: string
        402:
          description: not enough money
          schema:
              type: object
              properties:
                error:
                  type: string
```

В следующие часы, мистер Роджерс рассказал тебе все детали. В порядке для пересоздания сервера, ты можешь использовать эти характеристики для производства кучки кода которая сможет имплементировать backend часть. Это возможно для переписания всей части вручную, но в этом случае вор может скрыться от тебя пока ты это делаешь, так что тебе нужно будет сгенировать код как можно быстрее.

Каждый покупатель конфет одтает деньги, выбирает тип конфет для покупки и их количество. Эти данные будут отправлены на серевер через HTTP в формате JSON и затем:

1) Если сумма цен конфет меньше или равно количеству денег покупателя переданных в автомат, сервер отвечает с HTTP 201 и возвращает JSON с двумями полями "thanks": "Thank you!", и "change" (сдача), будет количество сдачи автоматы возвращенному обратно клиенту.
2) Если сумма больше чем количество денег переданное в автомат, сервер вернет HTTP 402 и ошибку с сообщением в JSON "Тебе нужно еще {amount} денег докинуть!", где {amount} это разница между ожидаемым и полученным. 
3) Если клиент предоставил негативное количество конфет или неправильный тип конфет (запомни - все пять типов конфет закодированы в две буквы, так вот все из них: `CE`, `AA`, `NT`, `DE`, `YR`, все остальные случаи недопустимы.) затем сервер отвечает с 400 и ошибкой внутри JSON что именно пошло не так. Ты можешь сделать это двумя различными способами, написать код для ручной проверки или модфицировать настройки сваггера выше для того чтобы покрыть данные случаи.

Запомни - вся информация с клиента и сервера должна быть в JSON, так ты можешь потестить свой сервер подобным образом:

```bash
curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:3333/buy_candy

{"change":5,"thanks":"Thank you!"}
```

или

```bash
curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/buy_candy

{"change":0,"thanks":"Thank you!"}
```

Так же тебе не нужно отслеживать наличие различных типов конфет самому, просто расчитывай что это будет сделано самим автоматом. Просто валидируй пользовательский ввод и вычисляй сдачу.

<h2 id="chapter-v" >Глава V</h2>
<h3 id="ex01">Упражнение 01: Закон и порядок</h3>

Ты откидываешься на спинку с улыбкой и чувством что видимо в этом случае все будет хорошо сделано. Мистер Роджерс тоже немного выдохнул. Но затем его лицо снова поменяло выражение.

"I know we've already paid for increased security at our datacenter" - he said a bit thoughtfully. - "...but what if this criminal desides to perform some [Man-in-the-middle](https://en.wikipedia.org/wiki/Man-in-the-middle_attack) trickery? My business will be destroyed again! People will lose their jobs abd I'll get bankrupt!"

"Я знаю мы уже оплатили за увеличение безопасности нашего датацентра" - сказал он немного вдумчиво. "...но что если этот преступник решит совершить атаку [Man-in-the-middle(человек посредине)](https://ru.wikipedia.org/wiki/%D0%90%D1%82%D0%B0%D0%BA%D0%B0_%D0%BF%D0%BE%D1%81%D1%80%D0%B5%D0%B4%D0%BD%D0%B8%D0%BA%D0%B0)? Мой бизнес будет снова уничтожен! Люди потеряют свои работы и  "

"Easy there, good sir" - you say with a smirk. - "I think I've got just what you need here."

So, you need to implement a certificate authentication for the server as well as a test client which will be able to query your API using a self-signed certificate and a local security authority to "verify" it on both sides.

You already have a server which supports TLS, but it is possible that you'll have to re-generate the code specifying an additional parameter, so it will be using use secure URLs.

Also, you'll need a local "certificate authority" to manage certificates. For our task [minica](https://github.com/jsha/minica) seems like a good enough solution. There is a link to a really helpful video in last Глава if you want to know more details about how Go works with secure connections.

So, because we're talking a full-blown mutual TLS authentication, you'll have to generate two cert/key pairs - one for the server and one for the client. Minica will also generate a CA file called `minica.pem` for you which you'll need to plug into your client somehow (your auto-generated server should already support specifying CA file as well as `key.pem` and `cert.pem` through command line parameters). Also, generating certificate may require you to use a domain instead of an IP address, so in examples below we will use "candy.tld".

Keep in mind, that because you're using a custom local CA you likely won't be able to query your API using cURL, web browser or tool like [Postman](https://www.postman.com/) anymore without tuning.

Your test client should support flags '-k' (accepts two-letter abbreviation for the candy type), '-c' (count of candy to buy) and '-m' (amount of money you "gave to machine"). So, the "buying request" should look like this:

```
~$ ./candy-client -k AA -c 2 -m 50
Thank you! Your change is 20
```

<h2 id="chapter-vi" >Глава VI</h2>
<h3 id="ex02">Упражнение 02: Старая корова</h3>

In a few days mister Rogers finally calls you with some great news - the thief was apprehended and immediately confessed! But candy businessman also had a small request.

"You seem like you really do know your way around machines, don't ya? There is one last thing I'd ask you to do, basically nothing. Our customers prefer something funny instead of just plain 'thank you', so my nephew Patrick wrote a program that generates some weird animal saying things. I think it's written in C, but that's not a problem for you, isn't it? Please don't change the code, Patrick is still improving it!"

Oh boy. You look through your emails and notice one from mister Rogers with a code attached to it:

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned int i;
unsigned int argscharcount = 0;

char *ask_cow(char phrase[]) {
  int phrase_len = strlen(phrase);
  char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
  strcpy(buf, " ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "_");
  }

  strcat(buf, "\n< ");
  strcat(buf, phrase);
  strcat(buf, " ");
  strcat(buf, ">\n ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "-");
  }
  strcat(buf, "\n");
  strcat(buf, "        \\   ^__^\n");
  strcat(buf, "         \\  (oo)\\_______\n");
  strcat(buf, "            (__)\\       )\\/\\\n");
  strcat(buf, "                ||----w |\n");
  strcat(buf, "                ||     ||\n");
  return buf;
}

int main(int argc, char *argv[]) {
  for (i = 1; i < argc; ++i) {
    argscharcount += (strlen(argv[i]) + 1);
  }
  argscharcount = argscharcount + 1;

  char *phrase = (char *)malloc(sizeof(char) * argscharcount);
  strcpy(phrase, argv[1]);

  for (i = 2; i < argc; ++i) {
    strcat(phrase, " ");
    strcat(phrase, argv[i]);
  }
  char *cow = ask_cow(phrase);
  printf("%s", cow);
  free(phrase);
  free(cow);
  return 0;
}
```

Looks like you'll have to return an ASCII-powered cow as a text in "thanks" field in response. When querying by cURL it will look like this:

```
~$ curl -s --key cert/client/key.pem --cert cert/client/cert.pem --cacert cert/minica.pem -XPOST -H "Content-Type: application/json" -d '{"candyType": "NT", "candyCount": 2, "money": 34}' "https://candy.tld:3333/buy_candy"
{"change":0,"thanks":" ____________\n< Thank you! >\n ------------\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||\n"}

```

Apparently, all you need is to reuse this `ask_cow()` C function without rewriting it in your Go code.

"Sometimes I think I have to drop this detective work and just go work as a Senior Engineer" - you grumble.

At least you should probably have as much candy as you want in return. Like, for the rest of your life.

<h2 id="chapter-vii" >Глава VII</h2>
<h3 id="reading">Reading</h3>

[Using the spec](https://goswagger.io/tutorial/custom-server.html)
[Secure connections](https://www.youtube.com/watch?v=kxKLYDLzuHA)
[Original cowsay](https://en.wikipedia.org/wiki/Cowsay)
