import { ApiService } from './api';

const BaseRoute: string = '/users';

export class UserService {
    public async getUser(id: number): Promise<User> {
        return ApiService.getData<User>(BaseRoute + '/' + id.toString());
    }

    public postUser(comic: User): Promise<number> {
        return ApiService.postData<User>(BaseRoute, comic);
    }

    public deleteUser(id: number): Promise<number> {
        return ApiService.deleteData<User>(BaseRoute + '/' + id.toString());
    }
}

export type User = {
    Id: number;
};
