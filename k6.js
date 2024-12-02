import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "5s", target: 20 },  
    { duration: "10s", target: 100 },
    { duration: "2s", target: 1 }, 
  ],
};

export default function () {
  const url = "http://127.0.0.1:8000/api/v1/books";

  const payload = JSON.stringify({
    name: "book",
    author: "sarath"
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const response = http.post(url, payload, params);

  check(response, {
    "status is 201": (r) => r.status === 201,
  });

}
