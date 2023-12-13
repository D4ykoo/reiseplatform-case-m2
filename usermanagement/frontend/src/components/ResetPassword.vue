<template>
  <div class="ml-auto mr-auto w-96">
    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">Username</span>
      </div>
      <input
        v-model="username"
        type="text"
        placeholder="Username"
        class="input input-bordered w-full"
      />
    </label>

    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">Old Password</span>
      </div>
      <input
        v-model="oldPassword"
        type="password"
        placeholder="Old Password"
        class="input input-bordered w-full"
      />
    </label>
    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">Old Password</span>
      </div>
      <input
        v-model="newPassword"
        type="password"
        placeholder="New Password"
        class="input input-bordered w-full"
      />
    </label>

    <div class="flex flex-row items-center justify-evenly align-middle">
      <button @click="back" class="btn btn-error btn-outline mt-6 w-2/5 flex">
        Cancel
      </button>
      <button @click="reset" class="btn btn-primary w-2/5 mt-6 hover:scale-105">
        Reset
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import type { ResetUser } from "@/models/UserModel";
import { LoginRegisterService } from "@/services/LoginRegisterService";
import { ref, type Ref } from "vue";

let loginRegisterService = new LoginRegisterService();

export default {
  name: "ResetPassword",
  setup() {
    const username = ref("");
    const oldPassword = ref("");
    const newPassword = ref("");

    return {
      username,
      oldPassword,
      newPassword,
    };
  },
  methods: {
    reset() {
      let resetUser: ResetUser = {
        oldLoginUser: {
          username: this.username,
          password: this.oldPassword,
        },
        newPassword: this.newPassword,
      };

      loginRegisterService
        .ResetPasswordRequest(resetUser)
        .subscribe((res: any) => {
        });
    },
    back() {
      this.$router.back();
    },
  },
};
</script>
