import http from 'k6/http';
import { check } from 'k6';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
    vus: 50,
    duration: '5s',
};

export default function () {
    const payload = JSON.stringify({
        Nome: `Pessoa ${uuidv4()}`,
        Apelido: `teste-${Math.random().toString()}`,
        Stack: ["Go", "C#", "Node", "Python"],
        Nascimento: "2002-03-12",
    });

    const headers = { 'Content-Type': 'application/json' };
    const res = http.post('http://localhost:8082/programadores', payload, { headers });

    check(res, {
        'status é 201 (Created)': (r) => r.status === 201,
    });
}