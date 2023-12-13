export type RequestUser = {
  username: string;
  firstname: string;
  lastname: string;
  email: string;
};

export type RegisterUser = {
  username: string;
  firstname: string;
  lastname: string;
  email: string;
  password: string;
};

export type LoginUser = {
  username: string;
  password: string;
};

export type UpdateUser = {
  requestUser: RequestUser;
  oldPassword: string;
};

export type User = {
  id: number;
  username: string;
  firstname: string;
  lastname: string;
  email: string;
};

export type ResetUser = {
  oldLoginUser: LoginUser;
  newPassword: string;
};
