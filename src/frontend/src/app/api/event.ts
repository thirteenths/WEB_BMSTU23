import { ApiService } from './api';

const BaseRoute: string = '/events';

export class EventService {
    public async getComic(id: number): Promise<Event> {
        return ApiService.getData<Event>(BaseRoute + '/' + id.toString());
    }

    public getAllComic(): Promise<Event[]> {
        return ApiService.getData<Event[]>(BaseRoute);
    }

    public postComic(comic: Event): Promise<number> {
        return ApiService.postData<Event>(BaseRoute, comic);
    }

    public deleteComic(id: number): Promise<number> {
        return ApiService.deleteData<Event>(BaseRoute + '/' + id.toString());
    }

    public putComic(comic: Event, id: number): Promise<number> {
        return ApiService.putData<Event>(BaseRoute + '/' + id.toString(), comic);
    }
}

export type Event = {
    Id: number;
};
