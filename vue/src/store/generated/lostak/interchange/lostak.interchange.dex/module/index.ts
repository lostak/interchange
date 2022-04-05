// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendCreatePair } from "./types/dex/tx";
import { MsgCancelBuyOrder } from "./types/dex/tx";
import { MsgSendBuyOrder } from "./types/dex/tx";
import { MsgSendSellOrder } from "./types/dex/tx";
import { MsgPort } from "./types/dex/tx";
import { MsgCancelSellOrder } from "./types/dex/tx";


const types = [
  ["/lostak.interchange.dex.MsgSendCreatePair", MsgSendCreatePair],
  ["/lostak.interchange.dex.MsgCancelBuyOrder", MsgCancelBuyOrder],
  ["/lostak.interchange.dex.MsgSendBuyOrder", MsgSendBuyOrder],
  ["/lostak.interchange.dex.MsgSendSellOrder", MsgSendSellOrder],
  ["/lostak.interchange.dex.MsgPort", MsgPort],
  ["/lostak.interchange.dex.MsgCancelSellOrder", MsgCancelSellOrder],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgSendCreatePair: (data: MsgSendCreatePair): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgSendCreatePair", value: MsgSendCreatePair.fromPartial( data ) }),
    msgCancelBuyOrder: (data: MsgCancelBuyOrder): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgCancelBuyOrder", value: MsgCancelBuyOrder.fromPartial( data ) }),
    msgSendBuyOrder: (data: MsgSendBuyOrder): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgSendBuyOrder", value: MsgSendBuyOrder.fromPartial( data ) }),
    msgSendSellOrder: (data: MsgSendSellOrder): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgSendSellOrder", value: MsgSendSellOrder.fromPartial( data ) }),
    msgPort: (data: MsgPort): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgPort", value: MsgPort.fromPartial( data ) }),
    msgCancelSellOrder: (data: MsgCancelSellOrder): EncodeObject => ({ typeUrl: "/lostak.interchange.dex.MsgCancelSellOrder", value: MsgCancelSellOrder.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
