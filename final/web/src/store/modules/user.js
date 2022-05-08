export default {
  state: {
    user: '' // 登录的用户
  },
  getters: {
    getUser(state) {
      return state.user
    }
  },
  mutations: {
    setUser(state, data) {
      state.user = data
    }
  },
  actions: {
    setUser({ commit }, data) {
      commit('setUser', data)
    }
  }
}
