import {beforeEach, describe, it, expect, jest} from '@jest/globals';
import { UserService, User } from './user';
import {ApiService} from "./api";

describe('Test : ComicService + getComic', () => {
    let userService: UserService;

    beforeEach(() => {
        userService = new UserService();
    });

    it('should return a Promise of Comic when calling getComic', () => {
        const mockUser: User = {Id: 1};
        jest.spyOn(ApiService, 'getData').mockResolvedValue(mockUser);

        // Call the getComic method
        const result = userService.getUser(1);

        // Expect the result to be the mock comic
        //expect(result).toEqual(mockUser);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + postComic', () => {
    let userService: UserService;
    const user: User = {Id: 1}

    beforeEach(() => {
        userService = new UserService();
    });

    it('should return a Promise of Comic when calling postComic', () => {
        jest.spyOn(ApiService, 'postData').mockResolvedValue(1);

        const result = userService.postUser(user);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + deleteComic', () => {
    let userService: UserService;

    beforeEach(() => {
        userService = new UserService();
    });

    it('should return a Promise of Comic when calling deleteComic', () => {
        jest.spyOn(ApiService, 'deleteData').mockResolvedValue(0);

        const result = userService.deleteUser(1);
        expect(result).toBeInstanceOf(Promise);
    });
});

