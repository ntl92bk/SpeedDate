﻿using SpeedDate.Network;
using SpeedDate.Network.Utils.IO;

namespace SpeedDate.Packets.Lobbies
{
    public class LobbyJoinTeamPacket : SerializablePacket
    {
        public int LobbyId;
        public string TeamName;

        public override void ToBinaryWriter(EndianBinaryWriter writer)
        {
            writer.Write(LobbyId);
            writer.Write(TeamName);
        }

        public override void FromBinaryReader(EndianBinaryReader reader)
        {
            LobbyId = reader.ReadInt32();
            TeamName = reader.ReadString();
        }
    }
}