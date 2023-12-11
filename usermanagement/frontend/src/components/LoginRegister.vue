<template>
  <div class="ml-auto mr-auto flex justify-center content-center align-middle">
    <div role="tablist" class="tabs tabs-boxed p-4 fix-width">
      <input
        type="radio"
        name="login-register"
        role="tab"
        class="tab m-auto"
        aria-label="Login"
      />
      <div role="tabpanel" class="tab-content p-10">
        <label class="form-control w-full ">
          <div class="label">
            <span class="label-text">Username</span>
          </div>
          <input
            type="text"
            placeholder="Username"
            class="input input-bordered w-full"
          />
        </label>
        <label class="form-control w-full">
          <div class="label">
            <span class="label-text">Password</span>
          </div>
          <input
            type="text"
            placeholder="Type here"
            class="input input-bordered w-full"
          />
          <div class="label">
            <span class="label-text-alt"></span>
            <span class="label-text-alt">Forgot password?</span>
          </div>
        </label>
        <button class="btn btn-primary w-full mt-4">Login</button>
      </div>

      <input
        type="radio"
        name="login-register"
        role="tab"
        class="tab"
        aria-label="Register"
        checked
      />
      <div role="tabpanel" class="tab-content p-10">
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
            type="text"
            placeholder="Password"
            class="input input-bordered w-full"
          />
        </label>
        <button class="btn btn-primary mt-4 w-9/12 flex ml-auto mr-auto">Register</button>
      </div>
    </div>
  </div>
</template>
