import axios from 'axios'

//读取商品概述的图片
const showInfoImgs = id =>
  axios.get(`/api/v1/info-imgs/${id}`).then(res => res.data)

//读取商品参数的图片
const showParamImgs = id =>
  axios.get(`/api/v1/param-imgs/${id}`).then(res => res.data)

export { showInfoImgs, showParamImgs }
