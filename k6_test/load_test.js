import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    thresholds: {
        http_req_failed: ['rate<0.01'], 
        http_req_duration: ['p(95)<1000'], // 95% of requests under 1s
    },
};

export default function () {
    const port = __ENV.API_PORT;
    const path = __ENV.API_PATH;
    
    const url = `http://localhost:${port}/${path}`; 

    const res = http.get(url);

    check(res, {
        'status is 200': (r) => r.status === 200,
        'has body': (r) => r.body.length > 0,
    });

    sleep(0.1); 
}