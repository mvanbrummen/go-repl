import axios from 'axios';

const BASE_URL = 'http://localhost:8080';

export function executeCode(language, code) {
    const url = `${BASE_URL}/code`;
    return axios.post(url, {
        language: language,
        code: code
    }, {
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(resp => resp.data)
        .catch(e => alert(e));
}

export function getVersion(language) {
    const url = `${BASE_URL}/version?lang=${language}`;
    return axios.get(url).then(resp => resp.data);
}