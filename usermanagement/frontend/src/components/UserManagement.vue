<template>
  <div
    class="overflow-x-auto overflow-y-auto top-0 flex flex-col place-items-start"
  >
    <div class="flex justify-end w-full">
      <button
        onclick="create_user_modal.showModal()"
        class="btn btn-primary btn-sm mr-2 hover:scale-105 ease-in-out"
      >
        Create
      </button>
      <CreateUserModal @eventCreateUser="reloadList()" />
    </div>
    <table class="table table-zebra">
      <!-- head -->
      <thead class="">
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
            <button @click="deleteUser(user.id)" class="hover:scale-105">
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
          <td class="w-3 h-3">
            <button
              onclick="edit_user_modal.showModal()"
              class="hover:scale-105"
            >
              <i class="fi fi-sr-pencil"></i>
            </button>
            <EditUserModal @eventUpdateUser="reloadList()" :parent-user-i-d="user.id" :parent-username="user.username" :parent-firstname="user.firstname" :parent-lastname="user.lastname" :parent-email="user.email"/>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import CreateUserModal from "./CreateUserModal.vue";
import EditUserModal from "./EditUserModal.vue";
import { UserManagementService } from "@/services/UserManagementService";
import type { User } from "@/models/UserModel";
import { ref, type Ref } from "vue";

let userManagementService = new UserManagementService();

export class Usermanagement {}

export default {
  name: "UserManagement",
  components: {
    CreateUserModal,
    EditUserModal,
  },
  setup() {
    let users: Ref = ref([]);

    return {
      users,
    };
  },
  mounted() {
    userManagementService.getAllUserRequests().subscribe((res: any) => {
      res.data.forEach((user: User) => {
        this.users.push(user);
      });
    });
  },
  methods: {
    reloadList() {
      this.users.length = 0
      userManagementService.getAllUserRequests().subscribe((res: any) => {
        res.data.forEach((user: User) => {
          this.users.push(user);
        });
      });
    },
    deleteUser(id: number) {
      userManagementService.deleteUser(id).subscribe((result: any) => {
        if (result.status === 200) {
          this.users = this.users.filter((elem: any) => {
            return elem.id !== id;
          });
        }
      });
    }
  },
};
</script>
