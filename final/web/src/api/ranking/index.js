import axios from 'axios'

// 读取热门家电
const listElecRanking = () =>
  axios.get('/api/v1/elec-rankings').then(res => res.data)

// 读取热门配件
const listAcceRanking = () =>
  axios.get('/api/v1/acce-rankings').then(res => res.data)

export { listElecRanking, listAcceRanking }
