import axios from 'axios'

// 获取数量
const showCount = id => axios.get(`/api/v1/counts/${id}`).then(res => res.data)

export { showCount }
