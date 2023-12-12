export type RequestUser = {
  username: string;
  firstname: string;
  lastname: string;
  email: string;
};

export type UpdateUser = {
    requestUser: RequestUser,
    oldPassword: string
}

export type User = {
  id: number;
  username: string;
  firstname: string;
  lastname: string;
  email: string;
};
