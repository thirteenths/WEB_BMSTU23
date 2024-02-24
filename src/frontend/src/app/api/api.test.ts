import axios from 'axios';
import { ApiService } from './api';
import {describe, expect, test, jest} from '@jest/globals';

describe('Test : Get', () => {

    test('should make a successful GET request and return data', async () => {
        const axiosGetMock1 = jest.spyOn(axios, 'get');
        axiosGetMock1.mockResolvedValue({
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });

        const response = await ApiService.getData('/route');
        expect(response).toEqual([
            { id: 1, name: 'John', email: 'john@gmail.com' },
        ]);
        expect(axiosGetMock1).toBeCalledWith('/route');
    });

    test('should handle error and return default data', async () => {
        const axiosGetMock2 = jest.spyOn(axios, 'get');
        axiosGetMock2.mockRejectedValue(new Error('Network Error'));

        const response = await ApiService.getData('/route');
        expect(response).toEqual({});
        expect(axiosGetMock2).toBeCalledWith('/route');
    });

    test('should return empty array when there is no data in response', async () => {
        const axiosGetMock3 = jest.spyOn(axios, 'get');
        axiosGetMock3.mockResolvedValue({ data: [] });
        const response = await ApiService.getData('/route');
        expect(response).toEqual([]);
        expect(axiosGetMock3).toBeCalledWith('/route');
    });
});

describe('Test : Post', () => {
    test('should make a successful POST request and return 1', async () => {
        const axiosGetMock1 = jest.spyOn(axios, 'post');
        axiosGetMock1.mockResolvedValue(0);

        const response = await ApiService.postData('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
        expect(response).toEqual(0);
        expect(axiosGetMock1).toBeCalledWith('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
    });

    test('should handle error and return default data', async () => {
        const axiosGetMock2 = jest.spyOn(axios, 'post');
        axiosGetMock2.mockRejectedValue(-1);

        const response = await ApiService.postData('/route', {});
        expect(response).toEqual(-1);
        expect(axiosGetMock2).toBeCalledWith('/route', {});
    });
});

describe('Test : Delete', () => {
    test('should make a successful Delete request and return 1', async () => {
        const axiosGetMock1 = jest.spyOn(axios, 'delete');
        axiosGetMock1.mockResolvedValue(0);

        const response = await ApiService.deleteData('/route');
        expect(response).toEqual(0);
        expect(axiosGetMock1).toBeCalledWith('/route');
    });

    test('should handle error and return default data', async () => {
        const axiosGetMock2 = jest.spyOn(axios, 'delete');
        axiosGetMock2.mockRejectedValue(-1);

        const response = await ApiService.deleteData('/route');
        expect(response).toEqual(-1);
        expect(axiosGetMock2).toBeCalledWith('/route');
    });
});

describe('Test : Put', () => {
    test('should make a successful Put request and return 1', async () => {
        const axiosGetMock1 = jest.spyOn(axios, 'put');
        axiosGetMock1.mockResolvedValue(0);

        const response = await ApiService.putData('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
        expect(response).toEqual(0);
        expect(axiosGetMock1).toBeCalledWith('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
    });

    test('should handle error and return default data', async () => {
        const axiosGetMock2 = jest.spyOn(axios, 'put');
        axiosGetMock2.mockRejectedValue(-1);

        const response = await ApiService.putData('/route', {});
        expect(response).toEqual(-1);
        expect(axiosGetMock2).toBeCalledWith('/route', {});
    });
});

describe('Test : Patch', () => {
    test('should make a successful Patch request and return 1', async () => {
        const axiosGetMock1 = jest.spyOn(axios, 'patch');
        axiosGetMock1.mockResolvedValue(0);

        const response = await ApiService.patchData('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
        expect(response).toEqual(0);
        expect(axiosGetMock1).toBeCalledWith('/route', {
            data: [{ id: 1, name: 'John', email: 'john@gmail.com' }],
        });
    });

    test('should handle error and return default data', async () => {
        const axiosGetMock2 = jest.spyOn(axios, 'patch');
        axiosGetMock2.mockRejectedValue(-1);

        const response = await ApiService.patchData('/route', {});
        expect(response).toEqual(-1);
        expect(axiosGetMock2).toBeCalledWith('/route', {});
    });
});
