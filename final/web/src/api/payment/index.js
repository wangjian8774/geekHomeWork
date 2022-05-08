import axios from 'axios'

// 初始化支付信息
const initPayment = form =>
  axios.post('/api/v1/payments', form).then(res => res.data)

export { initPayment }
