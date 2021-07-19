import { createStore } from 'vuex'



import topHotel from "./topHotel"
import login from "./login"
import bill from "./bill"

export default createStore({
  strict: true,
  modules: {
    topHotel,
    login,
    bill,
  }
})

