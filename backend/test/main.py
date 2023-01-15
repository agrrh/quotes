import requests
from faker import Faker

fake = Faker()

URL = "http://127.0.0.1:8081"


class TestIndex:
    def test_index(self):
        response = requests.get(URL)
        assert response.status_code == 201
        assert response.json().get("message") == "User created successfully."
        assert response.json().get("uuid")
        assert isinstance(response.json().get("uuid"), int)
        print(response.text)


# class TestIndex:
#     def test_index(self):
#         username = fake.email()
#         password = fake.password()
#         body = {"username": username, "password": password}
#         response = requests.post("http://127.0.0.1:8081/quotes/", json=body)
#         assert response.status_code == 201
#         assert response.json().get("message") == "User created successfully."
#         assert response.json().get("uuid")
#         assert isinstance(response.json().get("uuid"), int)
#         print(response.text)