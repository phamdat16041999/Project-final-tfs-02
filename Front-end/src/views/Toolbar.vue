<template>
  <div class="toolbar">
    <div class="container-xl">
      <div class="row">
        <nav class="navbar navbar-expand-lg col-12">
          <a class="navbar-brand toolbarText" @click="homePage">DKT Booking</a>
          <button
            class="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <div class="menu"></div>
            <div class="menu"></div>
            <div class="menu"></div>
          </button>

          <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto">
              <li class="nav-item active">
                <a class="nav-link toolbarText" @click="homePage"
                  >Home <span class="sr-only">(current)</span></a
                >
              </li>
              <li class="nav-item" v-if="login.role == 'HotelOwner'">
                <a class="nav-link toolbarText" href="/Hotelier">Hotel manage</a>
              </li>
              <li class="nav-item" v-if="login.login == true">
                <a class="nav-link toolbarText" href="/listBill">Your bill</a>
              </li>
              <li class="nav-item">
                <form class="form-inline my-2 my-lg-0">
                  <input
                    class="form-control mr-sm-2"
                    type="search"
                    placeholder="Search"
                    aria-label="Search"
                    v-model="searchData"
                  />
                  <button
                    class="btn btn-outline-success my-2 my-sm-0"
                    type="button"
                    @click="search"
                  >
                    Search
                  </button>
                </form>
              </li>
            </ul>
            <div class="nav-item" v-if="login.login == true">
             <i class='fab fa-facebook-messenger' style='font-size:30px;color:#2980b9;cursor: pointer;' @click="messenger"></i>
            </div>
            <div class="nav-item" v-if="login.login == false" style="margin-left:10px;">
              <a class="nav-link toolbarText" @click="loginPage">Login</a>
            </div>
            <div class="nav-item dropdown user" v-if="login.login" style="margin-left:10px;">
              <a
                class="nav-link dropdown-toggle toolbarText"
                href="#"
                id="navbarDropdown"
                role="button"
                data-toggle="dropdown"
                aria-haspopup="true"
                aria-expanded="false"
              >
                User
              </a>
              <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                <a class="dropdown-item" href="#">{{login.role}}</a>
                <a class="dropdown-item" href="#">View profile</a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item" href="#" @click="logOut">LogOut</a>
              </div>
            </div>
          </div>
        </nav>
      </div>
    </div>
  </div>
</template>
<script>
// import axios from "axios";
import { mapState } from "vuex";
export default {
  computed: mapState(["login"]),
  methods: {
    search(){
      this.$router.push("/search?name="+this.searchData);
    },
    homePage() {
      this.$router.push("/");
    },
    loginPage() {
      this.$router.push("/login");
    },
    messenger(){
      this.$router.push("/messenger");
    },
    logOut() {
      this.$store.dispatch("delUser");
      localStorage.removeItem("token");
      localStorage.removeItem("role");
      this.$router.push("/");
    },
  },
  created() {
    this.$store.dispatch("setUser");
  },
  data() {
    return {
      searchData: "",
    };
  },
};
</script>
<style scoped>
.toolbarText {
  color: white;
  cursor: pointer;
}
.user {
  width: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.toolbar {
  background-color: #474849;
  position: absolute;
  top:0;
  width: 100%;
  z-index: 1;
}
.menu {
  width: 35px;
  height: 5px;
  background-color: white;
  margin: 6px 0;
}
.nav-item{
   margin-left: 20px;
}
</style>