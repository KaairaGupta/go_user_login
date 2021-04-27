## USER REGISTRATION AND LOGIN

- This repository is made after following several tutorials on Go file structure, authentication, login, regsitration, and token generation.
- This API uses MongoDB to store the test user data
- It uses JWT for token generation
- User is granted access of a dummy API (dummy-api) on sucessful login.
- The app is hosted on http://localhost:8000 with routes as
	- /user/signup
	- /user/login
	- /dummy-api

## Screenshots Corresponding to Possible Situations

### Existing User

![Existing user error](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/1.png)

### Successful SignUp

![Successful Signup](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/2.png)

### Incorrect login attempts

![Wrong password](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/3.png)

![Wrong email](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/4.png)

### Successful Login

![Login successful](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/5.png)

### Incorrect token to access Dummy API

![Incorrect signature](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/6.png)

### Successful attempt to access Dummy API

![Successful attempt to access dummy api](https://github.com/KaairaGupta/go_user_login/blob/main/readme_images/7.png)