import axios, { AxiosResponse, AxiosError } from 'axios';


export class ApiService {
// Функция для отправки GET-запроса
    public static async getData<Type>(route: string): Promise<Type> {
        try {
            const response: AxiosResponse<Type> = await axios.get<Type>(route);
            const data: Type = response.data;
            console.log(data);
            return data;
        } catch (error) {
            const data = {} as Type;
            const axiosError = error as AxiosError;
            console.log(axiosError.message);
            return data;
        }
    }

// Функция для отправки POST-запроса
    public static async postData<Type>(
        route: string,
        data: Type,
    ): Promise<number> {
        try {
            const response: AxiosResponse<Type> = await axios.post<Type>(
                route,
                data,
            );
            const createdData: Type = response.data;
            console.log(createdData);
            return 0;
        } catch (error) {
            const axiosError = error as AxiosError;
            console.log(axiosError.message);
            return -1;
        }
    }

// Функция для отправки DELETE-запроса
    public static async deleteData<Type>(route: string): Promise<number> {
        try {
            const response: AxiosResponse<Type> = await axios.delete<Type>(route);
            const createdData: Type = response.data;
            console.log(createdData);
            return 0;
        } catch (error) {
            const axiosError = error as AxiosError;
            console.log(axiosError.message);
            return -1;
        }
    }

// Функция для отправки PUT-запроса
    public static async putData<Type>(
        route: string,
        data: Type,
    ): Promise<number> {
        try {
            const response: AxiosResponse<Type> = await axios.put<Type>(
                route,
                data,
            );
            const dataResponse: Type = response.data;
            console.log(dataResponse);
            return 0;
        } catch (error) {
            const axiosError = error as AxiosError;
            console.log(axiosError.message);
            return -1;
        }
    }

// Функция для отправки PATCH-запросов
    public static async patchData<Type>(
        route: string,
        data?: Type,
    ): Promise<number> {
        try {
            const response: AxiosResponse<Type> = await axios.patch<Type>(
                route,
                data,
            );
            const dataResponse: Type = response.data;
            console.log(dataResponse);
            return 0;
        } catch (error) {
            const axiosError = error as AxiosError;
            console.log(axiosError.message);
            return -1;
        }
    }
}