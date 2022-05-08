import axios from 'axios'

// 创建商品
const postProduct = form =>
  axios.post('/api/v1/products', form).then(res => res.data)

// 读商品详情
const showProduct = id =>
  axios.get(`/api/v1/products/${id}`).then(res => res.data)

// 读取商品列表
const listProducts = (category_id, start, limit) =>
  axios
    .get('/api/v1/products', { params: { category_id, start, limit } })
    .then(res => res.data)

//读取商品的图片
const showPictures = id => axios.get(`/api/v1/imgs/${id}`).then(res => res.data)

//搜索商品
const searchProducts = form =>
  axios.post('/api/v1/searches', form).then(res => res.data)

// 排行榜详情
const showRanking = () => axios.get('/api/v1/ranking/').then(res => res.data)

export {
  postProduct,
  showProduct,
  listProducts,
  showPictures,
  searchProducts,
  showRanking
}
