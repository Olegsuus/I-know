**NFS и Samba: Протоколы для Совместного Доступа к Файлам**

1) NFS — это протокол, разработанный для Linux и Unix-систем, который позволяет одному компьютеру делиться своими файлами с другими компьютерами по сети.


- Представьте, что у вас есть общий сетевой диск, доступный для всех компьютеров в вашей локальной сети. С помощью NFS вы можете монтировать (подключать) этот диск на своём компьютере и работать с файлами так, как будто они находятся на вашем локальном диске.



-	Пример использования: В офисе несколько сотрудников могут работать с одними и теми же файлами, хранящимися на сервере, используя NFS для доступа.


2) Samba

- Samba — это протокол, который позволяет компьютерам под управлением Windows и Linux делиться файлами и принтерами друг с другом.
	

- Если у вас есть компьютер с Windows и другой с Linux, Samba позволяет им обмениваться файлами, как если бы они находились в одной операционной системе.


- Пример использования: В смешанной сети, где работают пользователи Windows и Linux, Samba обеспечивает совместный доступ к общим папкам и принтерам.





**ClickHouse и Keycloak: Инструменты для Обработки Данных и Управления Доступом**

1) ClickHouse — это столбцовая база данных, разработанная для быстрого анализа больших объёмов данных.


- 	В отличие от традиционных реляционных баз данных, которые хранят данные по строкам, столбцовые базы данных хранят данные по столбцам. Это позволяет быстро выполнять аналитические запросы, например, подсчитывать средние значения по колонке.


Преимущества:
 
    Очень высокая скорость обработки больших объёмов данных.
	Эффективное использование памяти и дискового пространства.

Пример использования:

    Аналитические платформы, которые обрабатывают миллионы записей, такие как веб-аналитика или финансовые системы.




2) Keycloak

- Keycloak — это система управления идентификацией и доступом (IAM), которая помогает управлять пользователями и их доступом к приложениям и сервисам.


- Keycloak обеспечивает функции аутентификации (проверка личности пользователя) и авторизации (разрешение доступа к ресурсам). Он поддерживает вход через различные провайдеры, такие как Google, Facebook, а также корпоративные системы.


Преимущества:

	Упрощает управление пользователями и доступом.
	Поддерживает SSO (Single Sign-On), позволяя пользователям входить в несколько приложений одним набором учетных данных.


Пример использования:

	Компания использует Keycloak для управления доступом сотрудников к различным внутренним приложениям и сервисам.



**Столбцовые Базы Данных vs. Реляционные Базы Данных**

1) Реляционные Базы Данных (РБД)

- РБД, такие как MySQL, PostgreSQL и Oracle, хранят данные в таблицах, где каждая строка представляет собой запись, а столбцы — атрибуты записи.


Преимущества:

	Хорошо подходят для транзакционных систем (например, банковские системы).
	Поддерживают сложные связи между таблицами с помощью ключей.

Недостатки:

	Менее эффективны для аналитических запросов на больших объёмах данных.



2) Столбцовые Базы Данных


- Столбцовые БД, такие как ClickHouse, Cassandra и Google BigQuery, хранят данные по столбцам вместо строк.


Преимущества:

    Высокая скорость выполнения аналитических запросов, например, агрегирование данных по определённым столбцам.
    Эффективное сжатие данных, так как данные одного столбца часто похожи.


Недостатки:

    Менее подходят для систем, где важны быстрые операции вставки и обновления отдельных записей.




Когда использовать?

1) Реляционные БД:
	

    Когда важны сложные связи между данными и транзакционная целостность.
    Примеры: банковские системы, системы управления запасами.

2) Столбцовые БД:
       

    Для аналитики и обработки больших объёмов данных.
    Примеры: бизнес-аналитика, логирование, веб-аналитика.







**Hadoop и HDFS: Платформа для Обработки Больших Данных**

1) Hadoop

- 	Hadoop — это платформа с открытым исходным кодом, разработанная для хранения и обработки больших объёмов данных с помощью распределённых вычислений.


- Hadoop позволяет объединить несколько серверов в один кластер, где каждый сервер обрабатывает свою часть данных. Это позволяет масштабировать систему горизонтально, добавляя новые серверы по мере необходимости.


Компоненты Hadoop:

    HDFS (Hadoop Distributed File System): Система распределённого хранения данных.
    MapReduce: Модель программирования для обработки данных.



2) HDFS (Hadoop Distributed File System)

- HDFS — это файловая система, предназначенная для хранения больших объёмов данных, распределённых по множеству серверов.


1. Разделение данных:


    Большие файлы разбиваются на небольшие блоки (обычно по 128 МБ) и распределяются по различным серверам в кластере.
2. Репликация:

 
    Каждый блок данных хранится в нескольких копиях на разных серверах для обеспечения отказоустойчивости.
3. Доступ к данным:

 
    Когда пользователь или приложение запрашивает данные, HDFS определяет, на каких серверах хранятся нужные блоки, и собирает их вместе.



- Представьте, что у вас есть огромный файл журнала веб-сервера. HDFS разбивает этот файл на блоки и хранит их на разных серверах. Когда вам нужно проанализировать эти данные, Hadoop распределяет обработку блоков по всем серверам, что ускоряет процесс.




Преимущества Hadoop и HDFS

Масштабируемость:

    Лёгкость добавления новых серверов в кластер по мере роста объёма данных.
Отказоустойчивость:
    
    Благодаря репликации данные сохраняются даже при сбое некоторых серверов.
Высокая производительность:

    Распределённая обработка данных позволяет быстро выполнять аналитические задачи.



**Типы Данных в HDFS**

- Хранение данных в HDFS


Параметризация хранения:
        
    В HDFS данные хранятся в виде файлов, которые могут содержать различные типы данных, структурированные в определённом формате.


Популярные форматы файлов для HDFS:

    Text Files (Текстовые файлы): Простые текстовые данные.
    Sequence Files: Двоичные файлы, содержащие пары ключ-значение.
    Parquet: Колонковые файлы, оптимизированные для аналитических запросов.
    Avro: Двоичный формат данных, поддерживающий схемы.


- Выбор правильного формата данных влияет на эффективность хранения и скорость обработки данных. Например, колонковые форматы, такие как Parquet, лучше подходят для аналитических задач, так как позволяют быстро читать только нужные столбцы.



- Пример:

		Если вы храните лог-файлы веб-сервера в HDFS, вы можете использовать формат Parquet, чтобы позже быстро анализировать только столбцы с IP-адресами и временем запроса, игнорируя другие данные.