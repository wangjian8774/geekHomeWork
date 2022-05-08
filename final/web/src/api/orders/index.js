import axios from 'axios'

// 创建订单
const postOrder = form =>
  axios.post('/api/v1/orders', form).then(res => res.data)

// 读取订单
const listOrders = (user_id, type, start, limit) =>
  axios
    .get(`/api/v1/user/${user_id}/orders`, {
      params: { type, start, limit }
    })
    .then(res => res.data)

// 读取订单详情
const showOrder = order_id =>
  axios.get(`/api/v1/orders/${order_id}`).then(res => res.data)

export { postOrder, listOrders, showOrder }
