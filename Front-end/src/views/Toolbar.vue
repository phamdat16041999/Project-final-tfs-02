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
              <li class="nav-item">
                <a class="nav-link toolbarText" href="#">Link</a>
              </li>
              <li class="nav-item">
                <form class="form-inline my-2 my-lg-0">
                  <input
                    class="form-control mr-sm-2"
                    type="search"
                    placeholder="Search"
                    aria-label="Search"
                  />
                  <button
                    class="btn btn-outline-success my-2 my-sm-0"
                    type="submit"
                  >
                    Search
                  </button>
                </form>
              </li>
            </ul>
            <div class="nav-item" v-if="user == false">
              <a class="nav-link toolbarText" @click="login">Login</a>
            </div>
            <div class="nav-item dropdown user" v-if="user">
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
                <a class="dropdown-item" href="#">Action</a>
                <a class="dropdown-item" href="#">Another action</a>
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
export default {
  methods: {
    homePage() {
      this.$router.push("/");
    },
    login() {
      this.$router.push("/login");
    },
    logOut() {
      localStorage.removeItem("token");
      this.$router.push("/");
      this.user = false;
    },
  },
  created() {
    const cat = localStorage.getItem("token");
    if (cat != null) {
      this.user = true;
    } else {
      this.user = false;
    }
  },
  data() {
    return {
      user: false,
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
  border-radius: 46px;
  border-color: white;
  border-width: 2px;
  border-style: solid;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.toolbar {
  background-color: black;
  opacity: 0.7;
}
.menu{
    width: 35px;
  height: 5px;
  background-color: white;
  margin: 6px 0;
}
</style>