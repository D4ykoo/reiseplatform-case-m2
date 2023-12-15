import { defineStore } from "pinia";
import type { User } from "@/models/UserModel";

export const useUserStore = defineStore("user", {
  state: () => {
    return { user: {} as User };
  },
  getters: {
    receiveUser(): User {
      return this.user;
    },
  },
  actions: {
    changeUser(payload: User) {
      this.user = payload;
      console.log("changed", this.user);
    },
  },
});
