# Mobile-is-gone-authenticator
Fun use case: In case mobile phone is out of battery, but we need to have emergency code of our totp on our pc, this program will skip all security standards and ensure dominance of the user experience. Be careful with this program as it might make you lazy.

## Ideas that could be implemented
we could hold a list of many different TOTP's <br>

## How to use
Add a secrets/google_secret.txt and do 

    go run .

And then it should run. <br>
The reason why i called it google_secret.txt is because i tested it with google authenticator. <br>
Please note you should add your own secret, you can get it from a authenticator by exporting it. If you dont know how to do then contact me for help. <br>
When you get the code, it is probably something like 24 bytes, if that is the case you need to know it has to be at least 32 bytes. Therefore you will add = at the end of it until it reaches 32 bytes.

Example FE6JO4FEO3466AK634FEA <br>
lets just say that was 24 bytes, then you add the rest of the bytes like this until it reaches 32, test with len() function in go <br>
Example FE6JO4FEO3466AK634FEA====== <br>
And then when you have 32 bytes it should work <br>

<br>
<br>
Have fun!


