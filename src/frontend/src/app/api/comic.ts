import { ApiService } from './api';

const BaseRoute: string = '/comics';

export class ComicService {
    public async getComic(id: number): Promise<Comic> {
        return ApiService.getData<Comic>(BaseRoute + '/' + id.toString());
    }

    public getAllComic(): Promise<Comic[]> {
        return ApiService.getData<Comic[]>(BaseRoute);
    }

    public postComic(comic: Comic): Promise<number> {
        return ApiService.postData<Comic>(BaseRoute, comic);
    }

    public deleteComic(id: number): Promise<number> {
        return ApiService.deleteData<Comic>(BaseRoute + '/' + id.toString());
    }

    public putComic(comic: Comic, id: number): Promise<number> {
        return ApiService.putData<Comic>(BaseRoute + '/' + id.toString(), comic);
    }
}

export type Comic = {
    Id: number;
};
