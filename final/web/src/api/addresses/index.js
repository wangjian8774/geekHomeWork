import axios from 'axios'

// 创建购物车
const postAddress = form =>
  axios.post('/api/v1/addresses', form).then(res => res.data)

// 读取购物车
const showAddresses = user_id =>
  axios.get(`/api/v1/addresses/${user_id}`).then(res => res.data)

// 更新购物车
const updateAddress = form =>
  axios.put('/api/v1/addresses', form).then(res => res.data)

// 删除购物车
const deleteAddress = addressID =>
  axios
    .delete('/api/v1/addresses', {
      data: { address_id: addressID }
    })
    .then(res => res.data)

export { postAddress, showAddresses, updateAddress, deleteAddress }
