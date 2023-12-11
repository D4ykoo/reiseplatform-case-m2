<script setup lang="ts">
type User = {
  id: number;
  username: string;
  firstname: string;
  lastname: string;
  email: string;
};

let testuser: User = {
  id: 0,
  username: "testusername",
  firstname: "testUserFirst",
  lastname: "testUserLast",
  email: "test@tester.test",
};

let users: User[] = [testuser];

function deleteUser() {
  console.log("delete");
}
function editUser() {
  console.log("edit");
}
</script>

<template>
  <div
    class="overflow-x-auto overflow-y-auto top-0 flex place-items-start"
  >
    <table class="table table-zebra">
      <!-- head -->
      <thead>
        <tr>
          <th></th>
          <th>Username</th>
          <th>Firstname</th>
          <th>Lastname</th>
          <th>E-Mail</th>
          <th></th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user of users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.username }}</td>
          <td>{{ user.firstname }}</td>
          <td>{{ user.lastname }}</td>
          <td>{{ user.email }}</td>
          <td class="w-3 h-3">
            <button @click="deleteUser" class="hover:scale-105">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </td>
          <td>
            <button @click="editUser" class="hover:scale-105">
                <i class="fi fi-sr-pencil"></i>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
