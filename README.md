# lamports-signature

Lamport signature scheme is a one time signature scheme used to create a digital signature. It typically involves the use of a hash function, in this Lamport-esque implementation SHA-256 is used. This implementation is inspired by the problem set presented here: https://ocw.mit.edu/courses/mas-s62-cryptocurrency-engineering-and-design-spring-2018/pages/assignments/pset1-hash-based-signature-schemes/.

Lamport's signature scheme consists of three main steps:

## 1) Key Generation

Initially a private key is generated consisting of 2 rows of 256 numbers, where each number is 256 bits long. 

![image](https://user-images.githubusercontent.com/38185025/218060106-28387a48-c340-4129-b3b4-bf8336b3eeb8.png)

The private key is then used to generate the public key, this is done by hashing the each number of the private key to form public key also consisting of 2 rows of 256 numbers, where each number is 256 bits long:
![image](https://user-images.githubusercontent.com/38185025/218062761-dc06a70e-224f-4efa-849a-286b87a5a88f.png)


## 2) Signature generation

A signature is now generated using the message and the private key. First the messages SHA 256 hash is calculated then converted to a 256 bit binary representation (note: SHA256 output size is 256 bits): 

![image](https://user-images.githubusercontent.com/38185025/218064710-79d77223-e9a7-4186-991d-f65e11a6d44e.png)

The signature is now calculated by iterating throught the 256 bit binary representation and for every bit choosing a corresponding number from either row of the private key. If the bit is a 0 then row 1's number is chosen, if the bit is a 1 then row 2's number is chosen:
![image](https://user-images.githubusercontent.com/38185025/218068666-3fb1ad4e-2ba6-454e-92b9-0a7d960d5fe6.png)


## 3) Signature verifcation

In order to verify the signature, the signature, public key and message is required. The message recepient can follow a similar process as was followed in the signature generation step and can calculate the SHA 256 hash of the message and then obtain a 256 bit binary representation of the digest:

![image](https://user-images.githubusercontent.com/38185025/218064710-79d77223-e9a7-4186-991d-f65e11a6d44e.png)

The binary representation can then be used to select bits from the senders public key in a similar manner as was done in signature generation i.e. for every bit in the binary representation a corresponding number from either row of the public key is to be chosen, if the bit is a 0 then row 1's number is chosen, if the bit is a 1 then row 2's number is chosen. 
![image](https://user-images.githubusercontent.com/38185025/218072436-536ff1c6-0be1-46d4-af03-a6db56a00981.png)

Each number in the signature is then hashed and the output is compared to the sequence obtained above. If the signatures hash sequence matched the public key bits selected, the signature was successfully verified. 


