<template>
  <div>
      <div class="row">
          <div class="col-12" style="text-align:center;">
              <h3>Your hotel list</h3>
          </div>
          <div class="col-12" style="text-align: right;margin-bottom: 10px;}">
              <i class="material-icons" style="font-size:48px;color:#3498db; cursor: pointer;" @click="$emit('addHotel')">add_box</i>
          </div>
      </div>
    <table class="table table-striped table-bordered">
      <thead>
        <tr class="title">
          <th>Hotel name</th>
          <th>Address</th>
          <th>Description</th>
          <th>Update</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>John</td>
          <td>Doe</td>
          <td>john@example.com</td>
          <td class="icon"><i class="fa fa-edit" style="font-size:26px;color:#e74c3c"></i></td>
          <td class="icon"><i class="material-icons" style="font-size:30px;color:#e74c3c">delete</i></td>
        </tr>
        <tr>
          <td>Mary</td>
          <td>Moe</td>
          <td>mary@example.com</td>
          <td class="icon"><i class="fa fa-edit" style="font-size:26px;"></i></td>
          <td class="icon"><i class="material-icons" style="font-size:30px;">delete</i></td>
        </tr>
        <tr>
          <td>July</td>
          <td>Dooley</td>
          <td>july@example.com</td>
          <td class="icon"><i class="fa fa-edit" style="font-size:26px;"></i></td>
          <td class="icon"><i class="material-icons" style="font-size:30px;">delete</i></td>
        </tr>
      </tbody>
    </table>

  </div>
</template>
<script>
export default {
  setup() {},
};
</script>
<style scoped>
.icon{
    color:#e74c3c;
    cursor: pointer;
    text-align: center;
}
.icon:hover{
    color:#c0392b;
}
.title{
    text-align: center;
}
</style>