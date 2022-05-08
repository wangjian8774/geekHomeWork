import axios from 'axios'

// 公告详情
const showNotice = () => axios.get('/api/v1/notices').then(res => res.data)

export { showNotice }
