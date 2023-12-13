<!-- TODO: Props that fill the values of the inputs -->
<template>
  <dialog id="edit_user_modal" class="modal">
    <div class="modal-box">
      <h3 class="font-bold text-lg">Create User</h3>
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <div class="flex flex-row">
            <div class="flex flex-col mr-4">
              <label class="form-control w-full max-w-xs">
                <div class="label">
                  <span class="label-text">Username</span>
                </div>
                <input
                  v-model="newUsername"
                  type="text"
                  placeholder="Username"
                  class="input input-bordered w-full max-w-xs"
                />
              </label>
              <label class="form-control w-full max-w-xs">
                <div class="label">
                  <span class="label-text">E-Mail</span>
                </div>
                <input
                  v-model="newEmail"
                  type="text"
                  placeholder="E-Mail"
                  class="input input-bordered w-full max-w-xs"
                />
              </label>
            </div>
            <div class="flex flex-col">
              <label class="form-control w-full max-w-xs">
                <div class="label">
                  <span class="label-text">Firstname</span>
                </div>
                <input
                  v-model="newFirstname"
                  type="text"
                  placeholder="Firstname"
                  class="input input-bordered w-full max-w-xs"
                />
              </label>

              <label class="form-control w-full max-w-xs">
                <div class="label">
                  <span class="label-text">Lastname</span>
                </div>
                <input
                  v-model="newLastname"
                  type="text"
                  placeholder="Lastname"
                  class="input input-bordered w-full max-w-xs"
                />
              </label>
            </div>
          </div>

          <label class="form-control w-full">
            <div class="label">
              <span class="label-text">Old Password</span>
            </div>
            <input
              v-model="oldPassword"
              type="passsword"
              placeholder="Old Password"
              class="input input-bordered w-full"
            />
          </label>

          <label class="form-control w-full">
            <div class="label">
              <span class="label-text">New Password</span>
            </div>
            <input
              v-model="newPassword"
              type="passsword"
              placeholder="New Password"
              class="input input-bordered w-full"
            />
          </label>
          <div class="flex flex-row">
            <button class="btn btn-error btn-outline mt-6 w-2/5 flex ml-auto">
              Cancel
            </button>
            <button
              v-on:click="emitUpdate"
              class="btn btn-primary mt-6 flex ml-auto w-2/5 mr-auto"
            >
              Update
            </button>
          </div>
        </form>
      </div>
    </div>
  </dialog>
</template>

<script lang="ts">
import type { UpdateUser, User } from "@/models/UserModel";
import { UserManagementService } from "@/services/UserManagementService";
import { ref } from "vue";
import { useUserStore } from "@/assets/UserStore";
import AlertToasts from "./AlertToasts.vue";

let userManagementService = new UserManagementService();

export default {
  name: "EditUserModal",

  setup() {
    const store = useUserStore();
    // let storedUser : User= store.user

    const newUsername = ref(store.receiveUser.username);
    const newFirstname = ref(store.receiveUser.firstname);
    const newLastname = ref(store.$state.user.lastname);
    const newEmail = ref(store.$state.user.email);
    const oldPassword = ref("");
    const newPassword = ref("");
    const userID = ref(-1);

    store.$subscribe(
      () => {
        newUsername.value = store.receiveUser.username;
        newFirstname.value = store.receiveUser.firstname;
        newLastname.value = store.receiveUser.lastname;
        newEmail.value = store.receiveUser.email;
        userID.value = store.receiveUser.id;
      },
      { detached: true }
    );

    return {
      newUsername,
      newFirstname,
      newLastname,
      newEmail,
      newPassword,
      oldPassword,
      userID,
      store,
    };
  },
  methods: {
    emitUpdate() {
      let user: UpdateUser = {
        username: this.newUsername,
        firstname: this.newFirstname,
        lastname: this.newLastname,
        email: this.newEmail,
        newPassword: this.newPassword,
        oldPassword: this.oldPassword,
      };
      userManagementService
        .updateUser(this.userID, user)
        .subscribe((res: any) => {
          if (res.status === 200) {
            this.$emit("eventUpdateUser", true, `User ${user.username} edited`);            
          }
        });
    },
  },
};
</script>
