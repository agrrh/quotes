import os
import requests

from faker import Faker

Faker.seed(0)
fake = Faker()

URL = os.environ.get("APP_URL", "http://127.0.0.1:8081")


class TestIndex:
    def test_get_index(self: object) -> None:
        response = requests.get(URL)
        assert response.status_code == 200

        assert response.json().get("name") == "quotes"
        assert response.json().get("version")

    def test_post_quotes(self: object) -> None:
        author = fake.name()
        content = fake.paragraph(nb_sentences=3)
        body = {
            "author": author,
            "content": content,
        }

        response = requests.post(f"{URL}/quotes/", json=body)
        assert response.status_code == 200

        assert response.json().get("author") == author
        assert response.json().get("content") == content
        assert response.json().get("qid").isnumeric()

        # TODO Add negative cases

    def test_get_quotes(self: object) -> None:
        response = requests.get(f"{URL}/quotes/")
        assert response.status_code == 200

        quotes_list = response.json()
        assert type(quotes_list) is list
        assert len(quotes_list) > 0

    def test_get_quotes_single(self: object) -> None:
        qid = "1"
        response = requests.get(f"{URL}/quotes/{qid}")
        assert response.status_code == 200

        assert response.json().get("author")
        assert response.json().get("content")
        assert response.json().get("qid") == qid
