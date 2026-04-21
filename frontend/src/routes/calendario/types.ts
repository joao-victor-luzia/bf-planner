export interface CalendarTask {
    title: string;
    className: string;
    date: Date;
    len: number;
    isBottom?: boolean;
    detailHeader?: string;
    detailContent?: string;
    startCol?: number;
    startRow?: number;
    vlen?: number;
}

export interface CalendarDay {
    name: string;
    enabled: boolean;
    date: Date;
}