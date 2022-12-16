// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from "../wasmtypes";
import * as sc from "./index";

export class ImmutableFoundryCreateNewParams extends wasmtypes.ScProxy {
    tokenScheme(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamTokenScheme));
    }
}

export class MutableFoundryCreateNewParams extends wasmtypes.ScProxy {
    tokenScheme(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamTokenScheme));
    }
}

export class ImmutableFoundryDestroyParams extends wasmtypes.ScProxy {
    foundrySN(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }
}

export class MutableFoundryDestroyParams extends wasmtypes.ScProxy {
    foundrySN(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }
}

export class ImmutableFoundryModifySupplyParams extends wasmtypes.ScProxy {
    destroyTokens(): wasmtypes.ScImmutableBool {
        return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamDestroyTokens));
    }

    foundrySN(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }

    supplyDeltaAbs(): wasmtypes.ScImmutableBigInt {
        return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamSupplyDeltaAbs));
    }
}

export class MutableFoundryModifySupplyParams extends wasmtypes.ScProxy {
    destroyTokens(): wasmtypes.ScMutableBool {
        return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamDestroyTokens));
    }

    foundrySN(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }

    supplyDeltaAbs(): wasmtypes.ScMutableBigInt {
        return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamSupplyDeltaAbs));
    }
}

export class ImmutableHarvestParams extends wasmtypes.ScProxy {
    forceMinimumBaseTokens(): wasmtypes.ScImmutableUint64 {
        return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ParamForceMinimumBaseTokens));
    }
}

export class MutableHarvestParams extends wasmtypes.ScProxy {
    forceMinimumBaseTokens(): wasmtypes.ScMutableUint64 {
        return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ParamForceMinimumBaseTokens));
    }
}

export class ImmutableTransferAllowanceToParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }

    forceOpenAccount(): wasmtypes.ScImmutableBool {
        return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamForceOpenAccount));
    }
}

export class MutableTransferAllowanceToParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }

    forceOpenAccount(): wasmtypes.ScMutableBool {
        return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamForceOpenAccount));
    }
}

export class ImmutableAccountNFTsParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class MutableAccountNFTsParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class ImmutableBalanceParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class MutableBalanceParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class ImmutableBalanceBaseTokenParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class MutableBalanceBaseTokenParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class ImmutableBalanceNativeTokenParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }

    tokenID(): wasmtypes.ScImmutableTokenID {
        return new wasmtypes.ScImmutableTokenID(this.proxy.root(sc.ParamTokenID));
    }
}

export class MutableBalanceNativeTokenParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }

    tokenID(): wasmtypes.ScMutableTokenID {
        return new wasmtypes.ScMutableTokenID(this.proxy.root(sc.ParamTokenID));
    }
}

export class ImmutableFoundryOutputParams extends wasmtypes.ScProxy {
    foundrySN(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }
}

export class MutableFoundryOutputParams extends wasmtypes.ScProxy {
    foundrySN(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamFoundrySN));
    }
}

export class ImmutableGetAccountNonceParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class MutableGetAccountNonceParams extends wasmtypes.ScProxy {
    agentID(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
    }
}

export class ImmutableNftDataParams extends wasmtypes.ScProxy {
    nftID(): wasmtypes.ScImmutableNftID {
        return new wasmtypes.ScImmutableNftID(this.proxy.root(sc.ParamNftID));
    }
}

export class MutableNftDataParams extends wasmtypes.ScProxy {
    nftID(): wasmtypes.ScMutableNftID {
        return new wasmtypes.ScMutableNftID(this.proxy.root(sc.ParamNftID));
    }
}
