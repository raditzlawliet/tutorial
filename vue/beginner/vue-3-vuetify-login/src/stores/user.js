import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import api from '@/plugins/api'

export const useUserStore = defineStore('user', () => {
  // state
  const token = ref(localStorage.getItem('token') || null)
  const storedUser = ref(localStorage.getItem('user') || null)

  // getters
  const user = computed(() => {
    if (storedUser.value) {
      return JSON.parse(storedUser.value)
    }
    return storedUser.value
  })

  const isAuthorized = computed(() => !token.value)

  function storeToken(newToken) {
    // Save the token to localStorage
    localStorage.setItem('token', newToken)

    // Save the token and user to the store ktate
    token.value = newToken
  }

  function storeUser(newUser) {
    // Save the user to localStorage
    const stringifiedUser = JSON.stringify(newUser)
    localStorage.setItem('user', stringifiedUser)

    // Save the token and user to the store ktate
    storedUser.value = stringifiedUser
  }

  const login = async (credentials) => {
    // call API and await
    const response = await api.post('/auth/login', credentials)

    // if success, store token, otherwise throw err
    storeToken(response.data.data)

    // call API account/my
    const responseMy = await api.get('/account/my')

    // if success, store user, otherwise throw err
    storeUser(responseMy.data.data)
  }

  const logout = () => {
    // Remove token and user from localStorage
    localStorage.removeItem('token')
    localStorage.removeItem('user')

    // Remove from store
    token.value = null
    storedUser.value = null
  }

  return {
    token,
    user,

    isAuthorized,
    storeToken,
    storeUser,

    login,
    logout,
  }
})
