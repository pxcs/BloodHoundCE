import { render, screen } from 'src/test-utils';
import userEvent from '@testing-library/user-event';
import { rest, RequestHandler } from 'msw';
import { setupServer } from 'msw/node';
import EntityInfoDataTable from './EntityInfoDataTable';
import { EntityInfoPanelContextProvider } from './EntityInfoPanelContextProvider';
import { allSections } from './content';
import { ActiveDirectoryNodeKind } from 'bh-shared-ui';

const objectId = 'fake-object-id';
const sections = allSections[ActiveDirectoryNodeKind.GPO](objectId);

const queryCount = {
    controllers: {
        count: 0,
        limit: 128,
        skip: 0,
        data: [],
    },
    ous: {
        count: 8,
        limit: 128,
        skip: 0,
        data: [],
    },
    computers: {
        count: 3003,
        limit: 128,
        skip: 0,
        data: [],
    },
    users: {
        count: 1998,
        limit: 128,
        skip: 0,
        data: [],
    },
    'tier-zero': {
        count: 55,
        limit: 128,
        skip: 0,
        data: [],
    },
} as const;

const handlers: Array<RequestHandler> = [
    rest.get(`api/v2/gpos/${objectId}/:asset`, (req, res, ctx) => {
        const asset = req.params.asset as keyof typeof queryCount;
        return res(ctx.json(queryCount[asset]));
    }),
];

const server = setupServer(...handlers);

describe('EntityInfoDataTable', () => {
    describe('Node count', () => {
        beforeAll(() => server.listen());
        afterEach(() => server.resetHandlers());
        afterAll(() => server.close());

        it('sums nested section node counts', async () => {
            render(
                <EntityInfoPanelContextProvider>
                    <EntityInfoDataTable {...sections[0]} />
                </EntityInfoPanelContextProvider>
            );
            const sum = await screen.findByText('5,064');
            expect(sum).not.toBeNull();
        });

        it('displays ! icon when one of the Affected Object calls fail', async () => {
            console.error = vi.fn();
            server.use(
                rest.get(`api/v2/gpos/${objectId}/ous`, (req, res, ctx) => {
                    return res(ctx.status(500));
                })
            );

            render(
                <EntityInfoPanelContextProvider>
                    <EntityInfoDataTable {...sections[0]} />
                </EntityInfoPanelContextProvider>
            );

            const errorIcon = await screen.findByTestId('ErrorOutlineIcon');

            expect(errorIcon).not.toBeNull();
        });

        it('displays 0 when a given sections returns empty, and sums the rest of the sections correctly', async () => {
            server.use(
                rest.get(`api/v2/gpos/${objectId}/ous`, (req, res, ctx) => {
                    const _ous = { ...queryCount.ous, count: undefined };
                    return res(ctx.json(_ous));
                })
            );

            const user = userEvent.setup();
            render(
                <EntityInfoPanelContextProvider>
                    <EntityInfoDataTable {...sections[0]} />
                </EntityInfoPanelContextProvider>
            );

            const sum = await screen.findAllByText('5,056');
            expect(sum).not.toBeNull();

            const button = await screen.findByRole('button');
            await user.click(button);

            const zero = await screen.findByText('0');
            expect(zero.textContent).toBe('0');
        });
    });
});
