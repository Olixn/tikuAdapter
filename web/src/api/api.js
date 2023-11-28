import axios from 'axios'

const instance = axios.create({
    baseURL: '/adapter-service',
    timeout: 30 * 1000
})


export async function getQuestions(params){
    return await instance.get('/questions', {params})
}

export async function updateQuestions(id,data) {
    return await instance.put(`/questions/${id}`,data)
}

export async function addQuestions(data) {
    return await instance.post(`/questions`,data)
}

export async function delQuestions(id) {
    return await instance.delete(`/questions/${id}`)
}



