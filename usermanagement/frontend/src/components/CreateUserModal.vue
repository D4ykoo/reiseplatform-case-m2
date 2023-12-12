<template>
    <dialog id="create_user_modal" class="modal">
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
                type="text"
                placeholder="Lastname"
                class="input input-bordered w-full max-w-xs"
              />
            </label>
          </div>
        </div>

        <label class="form-control w-full">
          <div class="label">
            <span class="label-text">Password</span>
          </div>
          <input
            type="passsword"
            placeholder="Password"
            class="input input-bordered w-full"
          />
        </label>
        <button class="btn btn-primary mt-4 w-9/12 flex ml-auto mr-auto">
          Register
        </button>
        </form>
        </div>
    </div>
    </dialog>
</template>

<script lang="ts">
export default {
    name: "CreateUserModal", 
}
</script>