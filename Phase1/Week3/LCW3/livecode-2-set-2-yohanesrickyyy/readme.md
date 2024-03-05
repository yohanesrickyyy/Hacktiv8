# Live-Code-2-Phase-1

## RULES

1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file Go (\*.go) yang telah disediakan.
5. Waktu pengerjaan: **120 menit**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. Lakukan `git push origin <branch-name>` dan create pull request **hanya jika waktu Live Code telah usai (bukan ketika kalian sudah selesai
mengerjakan)**. Tuliskan nama lengkap kalian saat membuat pull request dan assign buddy.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.


### KETENTUAN

Here are some dos and don'ts to consider when working on this livecode:

Dos:

- Do read and understand the problem statement and requirements carefully before starting to code.
- Do break down the problem into smaller, manageable parts and tackle each one individually.
- Do test your code frequently and thoroughly to ensure that it is functioning as intended.
- Do use good programming practices, such as meaningful variable names, clear comments, and properly formatted code.
- Do ask for help if you get stuck or need clarification on a specific concept or feature.

Don'ts:

- Don't rush through the problem or try to solve it all at once.
- Don't copy and paste code from external sources without fully understanding how it works.
- Don't hardcode values or rely on assumptions that may not hold true in all cases.
- Don't forget to handle error cases or edge cases, such as invalid input or unexpected behavior.
- Don't hesitate to refactor your code or make improvements based on feedback or new insights.

### NUMBER 1 LIVE CODE 2

## **Hacktiv8 Order System**

### Restrictions

Do:

- Use channels to communicate between goroutines.
- Use the sync.WaitGroup to wait for all the orders to be served before the program exits.
- Use descriptive variable names to make the code easier to understand.
- Use comments and documentation to explain the purpose of each function and channel, as well as how to run the program and what it does.
- Test the code with different numbers of orders to make sure it can handle a variable workload.
- Handle errors and exceptions properly to prevent the program from crashing.

Don't:

- Use global variables to communicate between goroutines.
- Use sleep statements to slow down the program unnecessarily.
- Use non-descriptive variable names that make the code harder to understand.
- Neglect to add comments or documentation to the code.
- Hard-code values that could change, such as the number of orders to be processed.
- Ignore errors or exceptions and let the program crash.

### Objectives

- Mampu menyelesaikan masalah yang diberikan.
- Memahami konsep dan cara penggunaan `looping`.
- Mengerti konsep concurrent
- Mengerti konsep `sync.WaitGroup` 
- Mengerti konsep dan cara penggunaan `go routine` dan `channel`.

#### Directions

You have been hired by a large logistics company to develop a new shipment tracking system that can handle a high volume of shipments and provide a better customer experience. The company currently uses a manual system for tracking shipments and managing inventory, which has resulted in delayed shipments and errors in item tracking.

Your task is to create a software system that can automate the shipment tracking process and enhance efficiency in inventory management. This system should allow customers to monitor the status of their shipments through a user interface, and it should automatically update the shipment statuses in the warehouse and during transit. Once the shipments have reached their destination, the system should automatically notify the customers through notifications.

To ensure that the system is scalable and can handle a high volume of shipments, it should be designed to run concurrently, with multiple shipments being processed and monitored simultaneously. The system should also be robust and able to handle errors and exceptions gracefully, to prevent the program from crashing or losing data. You will need to use Golang to create this software system, and you should follow best practices for coding style, documentation, and testing. The system should be designed to be modular and extensible so that it can be easily modified or expanded in the future as the company's needs change.

You are tasked with creating a logistics application that handles three main features:

- Receiving shipment information from customers.
- Processing and tracking shipments in the warehouse.
- Notifying customers of shipment statuses.

Your application should run concurrently to handle multiple shipments at the same time, and it should be built in Golang.

### Expected

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V3-LC2/blob/main/Hasil.gif)