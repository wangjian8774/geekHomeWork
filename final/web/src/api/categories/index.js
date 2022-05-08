import axios from 'axios'

// 创建分类
const postCategory = form =>
  axios.post('/api/v1/categories', form).then(res => res.data)

// 读取分类列表
const listCategories = () =>
  axios.get('/api/v1/categories').then(res => res.data)

export { postCategory, listCategories }
